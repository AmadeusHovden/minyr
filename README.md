### MINYR
## må gjøres:

# Lærer må kunne klone repositorien, kompilere 'main.go' filen med kommandoen 'go build' og utføre programmet med kommandoen 'minyr' som gir disse tre valgene: convert, average, exit. Lærer skal bruke en Linux terminal for å evaluere besvarelsen.??

# Alt arbeid skal gjennomføres i instansen av Linux operativsystemet (Ubuntu) i en Docker kontainer, som ble installert og konfigurert på Seminar II  (unntaket er hvis studenten har Linux OS installert direkte på datamaskinen, dvs. som vertssystem).

# endre navn på den nye fila til: kjevik-temp-fahr-20220318-20230318.csv  !!!!!


## TESTER
## Tester som skal være implementert:

# gitt "Kjevik;SN39040;18.03.2022 01:50;6" ønsker å få (want) "Kjevik;SN39040;18.03.2022 01:50;42.8"

# gitt "Kjevik;SN39040;07.03.2023 18:20;0" ønsker å få (want) "Kjevik;SN39040;07.03.2023 18:20;32.0"

# gitt "Kjevik;SN39040;08.03.2023 02:20;-11" ønsker å få (want) "Kjevik;SN39040;08.03.2023 02:20;12.2"

# gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want) "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen
