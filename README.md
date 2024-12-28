
# 1. KONTEKST: SYMULACJA\
Rozważmy następującą symulację: w dwuwymiarowym obszarze n x m (wymiary w metrach) porusza się w dowolnym kierunku oraz z losową szybkością (nie większą niż 2.5[]) i zdrowych osobników. Szybkość oraz kierunek ich przemieszczania może ulegać losowym zmianom w czasie (nie może jednak przekraczać wyżej ustalonego górnego limitu). W momencie dotarcia do dowolnej granicy obszaru każdy osobnik może:\
zawrócić do wewnątrz obszaru (prawdopodobieństwo 50 procent) opuścić obszar (prawdopodobieństwo 50 procent)\
W trakcie trwania symulacji nowi osobnicy wkraczają do obszaru w losowych punktach na jego granicach (częstotliwość wkraczania oraz początkową liczebność i dobrać tak by zachować ciągłość populacji). Dla każdego wkraczającego osobnika istnieje prawdopodobieństwo zakaże- nia wirusem wynoszące 10 procent.\
Każdy osobnik w populacji jest:\
odporny na zakażenie\
albo (alternatywa wykluczająca) wrażliwy na zakażenie\
jeżeli osobnik jest wrażliwy na zakażenie to jest:\
zdrowy\
albo (alternatywa wykluczająca)\
zakażony\
jeżeli osobnik jest zakażony to:\
posiada objawy\
albo (alternatywa wykluczająca)\
nie posiada objawów\
Osobnik zdrowy oraz nieodporny na zakażenie zaraża się od osobnika zakażonego wtedy i tylko wtedy gdy: a) odległość między nimi nie przekracza 2[m] oraz (koniunkcja) b) czas gdy ta odległość jest utrzymana wynosi nie mniej niż 3[s] symulacji. Prawdopodobieństwo zakażenia się od osobnika bezobjawów wynosi 50 procent, a od osobnika z objawowym przejściem choroby 100 procent. Zakażony osobnik podtrzymuje zakażenie od 20 do 30 sekund symulacji po czym zdrowieje, uzyskując odporność.
# 2. ZADANIE LABORATORYJNE\
Zaprojektować oraz zaimplementować rozwiązanie symulujące rozwój zakażenia w populacji. Umożliwić zapis oraz wczytywanie stanu symulacji w dowolnym momencie t od rozpoczęcia. Rozwiązanie ma umożliwiać wizualizację przemieszczania się oraz zakażania. Rozważyć dwa przypadki:
początkowa populacja oraz losowane osobniki nie posiadają odporności
część początkowej populacji oraz wylosowanych osobników posiada odporność Wykorzystać wektory z Laboratorium nr 2 w celu modelowania ruchu osobników.\
UWAGI\
Każda sekunda symulacji winna składać się z 25 kroków.
