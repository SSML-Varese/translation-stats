# translation-stats

Whenever one of students or group of students translate an article for an online newspaper it has to be revised by someone. The date of the translation, the information of the student(s) and the editor is stored in the filename. At the end of the academic year, there is a need to count how many translations each student has done.

This script will generate statistics for a collection of files where the filename contains the date, the author(s), and the editor.

The basic filenaming structure is: `DATE - AUTHOR(S)-EDITOR`

The authors can be written as follows:

- `Author`
- `Author & Author`
- `Author, Author & Author`

If there are multiple files with the same date and same author, the author can be written as:

- `Author 1`
- `Author 2`

Currently, the script ignores the date and editor information.
