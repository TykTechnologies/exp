#!/usr/bin/env python3
import argparse
import re
from collections import defaultdict

def find_headers(lines):
    """
    Returns a list of (line_index, level, title) tuples for all H2/H3 headers.
    """
    header_pattern = re.compile(r'^(#{2,3})\s+(.*)$')
    headers = []
    for i, line in enumerate(lines):
        match = header_pattern.match(line.strip())
        if match:
            level, title = match.groups()
            headers.append((i, len(level), title.strip()))
    return headers


def log_repeated_headers(headers):
    """
    Prints a summary of repeated H2/H3 headers.
    """
    title_map = defaultdict(list)
    for i, _, title in headers:
        title_map[title].append(i + 1)  # convert to 1-based line numbers

    repeated = {h: lines for h, lines in title_map.items() if len(lines) > 1}

    if not repeated:
        print("✅ No repeated H2 or H3 headers found.\n")
        return set()

    print("⚠️ Repeated H2/H3 Headers Found:\n")
    for header, lines in repeated.items():
        print(f'"{header}" — {len(lines)} times (lines: {", ".join(map(str, lines))})')
    print()

    # Return lowercase versions for easier comparison
    return {h.lower() for h in repeated.keys()}


def remove_duplicate_sections(lines, headers, duplicates):
    """
    Removes duplicate sections, keeping only the first occurrence.
    """
    seen = set()
    to_remove = set()

    for idx, (_, _, title) in enumerate(headers):
        t_lower = title.lower()
        if t_lower in duplicates:
            if t_lower in seen:
                start = headers[idx][0]
                end = headers[idx + 1][0] if idx + 1 < len(headers) else len(lines)
                for i in range(start, end):
                    to_remove.add(i)
            else:
                seen.add(t_lower)

    return [line for i, line in enumerate(lines) if i not in to_remove]


def main():
    parser = argparse.ArgumentParser(
        description="Detect and remove duplicate H2/H3 headers and their sections from a Markdown file."
    )
    parser.add_argument("--input-file", required=True, help="Path to the markdown (.md) file")
    args = parser.parse_args()

    with open(args.input_file, "r", encoding="utf-8") as f:
        lines = f.readlines()

    headers = find_headers(lines)
    duplicates = log_repeated_headers(headers)
    cleaned_lines = remove_duplicate_sections(lines, headers, duplicates)

    with open(args.input_file, "w", encoding="utf-8") as f:
        f.writelines(cleaned_lines)

    print(f"✅ Deduplication complete. File updated: {args.input_file}")


if __name__ == "__main__":
    main()
