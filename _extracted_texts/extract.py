from pathlib import Path
import zipfile
import re
from pypdf import PdfReader

base = Path('/work')
out = base / '_extracted_texts'
out.mkdir(exist_ok=True)
files = list((base / 'Go').glob('*.pdf')) + list((base / 'Go').glob('*.docx')) + list((base / 'ПКП').glob('*.pdf')) + list((base / 'ПКП').glob('*.docx')) + list((base / 'РХД').glob('*.pdf')) + list((base / 'РХД').glob('*.docx'))

for path in files:
    text = ''
    if path.suffix.lower() == '.pdf':
        reader = PdfReader(str(path))
        parts = []
        for page in reader.pages:
            parts.append(page.extract_text() or '')
        text = '\n'.join(parts)
    elif path.suffix.lower() == '.docx':
        with zipfile.ZipFile(path) as zf:
            xml = zf.read('word/document.xml').decode('utf-8', errors='ignore')
        text = re.sub(r'</w:p>', '\n', xml)
        text = re.sub(r'<[^>]+>', '', text)
    (out / (path.name + '.txt')).write_text(text, encoding='utf-8')
    print(path.name)
