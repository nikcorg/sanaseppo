# Sanaseppo

Sanaseppo on apulainen [Sanaseppä-sanapelille](https://sanaseppa.com), samassa hengessä kuin NYT:n Spelling Bee -pelin vihjesivu.

## Ohjelman kääntäminen

Ohjelman kääntäminen edellyttää `make` ja `go` komentojen asentamista. Käännä ohjelma komentamalla `make`.

Sanalistan voi päivittää komennolla `make wordlist`. Huom! Lähdetiedostoa ei noudeta, vaan sen oletetaan löytyvän levyltä. Mikäli lähdetiedoston nimi muuttuu (kirjoitushetkellä nykysuomensanalist2024.txt), päivitä nimitieto `Makefile` tiedostoon `WORDLIST_SOURCE` muuttujaan.

Sanalistan päivittäminen vaatii ohjelman kääntämisen uudelleen.

## Ohjelman käyttäminen

Sanaseppo on komentoriviltä käytettävä ohjelma. Anna argumenteiksi keskikirjain ja loput kirjaimet. Tuloste näyttää listan kaksikirjaimisia etuliitteitä, jonka perässä on etuliitteelle löytyneiden sanojen kokonaislukumäärä, sekä löytyneiden sanojen lukumäärä sanan pituuden mukaan.

Yhteenvetorivi tulosteen lopussa kertoo kaikkien löytyneiden sanojen lukumäärän, pangramien määrän ja täydellisten pangramien määrän.

```
% ./bin/seppo h spnuoe
he:  3  4=1     5=1     6=1
ho:  3  5=3
hu: 11  4=2     5=5     7=2     8=1     9=1
nu:  1  7=1
oh:  6  4=1     5=3     6=1     7=1
pe:  1  9=1
pu:  3  4=2     10=1
sh:  1  4=1
su:  1  7=1
un:  1  4=1

31 words (2 pangrams, 0 perfect)
```

Antamalla kolmanneksi ohjelma-argumentiksi etuliitteen, seppo paljastaa löytyneet sanat.

## Lisenssi

Sanaseppo on [Unlicense](https://unlicense.org) lisenssin alainen.

Lähdekoodi sisältää myös Suomen Kielitoimiston julkaiseman [Nykysuomen sanalistan](https://kotus.fi/sanakirjat/kielitoimiston-sanakirja/nykysuomen-sana-aineistot/nykysuomen-sanalista/), joka on [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/deed.fi) lisenssin alainen.