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

## Aufgaben

### Aufgabe 1

Im Package `tasks/circularbuffer` ist ein Typ `CircularBuffer` definiert,
die Methoden sind aber noch nicht vollständig implementiert.
Vervollständigen Sie dieses Package, so dass die Tests in `tasks/circularbuffer/buffer_test.go`
und `tasks/circularbuffer/buffer_empty_test.go` erfolgreich durchlaufen.

*Anmerkung:* In den Verzeichnissen `hints` und `solutions`
finden Sie Hinweise und eine mögliche Lösung.

### Aufgabe 2

Die Implementierung in `circularbuffer/buffer.go` erlaubt es, beliebige Typen
im Puffer miteinander zu mischen.
Dies wird auch in der Datei `buffer_different_types_test.go` im Lösungs-Package demonstriert.

Verwenden Sie Generics, um den Typ des Puffers zu parametrisieren,
so dass nur Elemente eines bestimmten Typs im Puffer gespeichert werden können.
Passen Sie die Tests in `circularbuffer/buffer_test.go` und `circularbuffer/buffer_empty_test.go`
entsprechend an.

*Anmerkung:* Im den Verzeichniss `tasks/circularbuffer_generic` finden Sie einen
Anfang, bei dem die Buffer-Impelementierung aus der Lösung zu Aufgabe 1
zusammen mit einem Test vorgegeben ist, der zeigt, wie es am Ende funktionieren sollte.
Diese Vorgabe compiliert aber nicht. Ihre Aufgabe ist es, die Buffer-Implementierung
anzupassen. Alternativ zur Vorgabe können Sie auch Ihre Lösung zu Aufgabe 1
verwenden.
