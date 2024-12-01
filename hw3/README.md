# Uniq - утилита, позволяющая определять уникальные следующие друг за другом строки с возможностью задания параметров определения и вывода уникальных строк

## Параметры

`-с` - подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.

`-d` - вывести только те строки, которые повторились во входных данных.

`-u` - вывести только те строки, которые не повторились во входных данных.

`-f num_fields` - не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.

`-s num_chars` - не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).

`-i` - не учитывать регистр букв.

## Использование

`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`
