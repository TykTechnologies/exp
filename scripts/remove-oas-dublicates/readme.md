# 🧹 Markdown Header Deduplicator

A small Python utility that scans a Markdown file for **repeated H2 (`##`) and H3 (`###`) headers**, prints their occurrences, and **removes duplicate sections**, keeping only the first instance of each header.


## 🧭 Example

### Input File (`example.md`)

```markdown
## Title

Some intro content.

### Hello

First section

### Hello

Duplicate section content

### Nice

Another section
```

### Output File (after running the script)

```markdown
## Title

Some intro content.

### Hello

First section

### Nice

Another section
```

### Console Output

```bash
⚠️ Repeated H2/H3 Headers Found:

"**Hello**" — 2 times (lines: 5, 9)

✅ Deduplication complete. File updated: example.md
```

---

## 🧰 Usage

### 1. Run the script

```bash
python3 dedupe_headers.py --input-file path/to/your-file.md
```

---
