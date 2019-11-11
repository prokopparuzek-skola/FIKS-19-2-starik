---
header-includes:
	\usepackage[czech]{babel}
	\usepackage{a4wide}
---
# Důkaz

## Řešení

Provedl jsem úvahu a došel k těmto závěrům. Pokud je počet cihel na patře sudý, vždy vyhraje stařík, protože z celého 
sloupku lze bez jeho spadnutí odebrat až sudý počet cihel. Z každého patra můžu vzít až polovinu cihel, ta je sudá, 
třeba všechno až na 2 cihly uprostřed. Vzhledem k tomu, že stařík si bere každou druhou cihlu, nutně odebere tu 
poslední, kdy sloupek ještě nespadne a na mě zbyde ta při které sloup už spadne. Pokud je počet cihel na patře lichý, 
závisí to na počtu pater, pokud je počet pater sudý, též vyhraje stařík. Jeho cílem je aby se celkem odebral sudý počet 
cihel, předtím než sloupek spadne, takže pokud já vezmu středovou cihlu, čímž by se počet odebraných cihel před 
spadnutím změnil na lichý, vezme se vše až na například dvě krajní cihly, a vyhrál bych, on vezme též středovou a změní 
to zpět na sudý. Pokud je počet pater lichý vyhraji já, protože pokud na začátku budu brát pouze středové díly, než 
dojdou. Tak bez ohledu na to co odebere stařík bude počet cihel odebraný před spadnutím lichý. A proto poslední cihlu 
před spadnutím odeberu já.

### Složitost
Složitost odhadu výjry je konstantní, takže celková složitost bude záviset pouze na počtu úloh k řešení a bude to: O(`N`).
