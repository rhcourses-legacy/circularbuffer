# Beispielimplementierung und Übungsaufgaben zu einem Ringpuffer

## Beschreibung

Ein Ringpuffer ist eine listenartige Datenstruktur mit fester maximaler Größe,
in die am einen Ende Daten hineingeschrieben und am anderen Ende wieder
ausgelesen werden können.  Wenn der Puffer voll ist, wird beim Schreiben das
älteste Element überschrieben.

Man kann sich einen Ringpuffer als einen Ring mit zwei Zeigern vorstellen.
Der eine Zeiger zeigt auf das älteste Element,
der andere auf die Position hinter dem jüngsten Element.
Beim Schreiben wird der jüngste Zeiger um eins weitergeschoben, beim Lesen der
älteste. Wenn die Zeiger aufeinandertreffen, ist der Puffer voll.
Implementiert wird der Ringpuffer als Array, das mit den Zeigern indiziert wird.
Beim Weiterschieben der Zeiger wird der Index modulo der Puffergröße genommen.
