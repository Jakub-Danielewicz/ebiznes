Jakub Danielewicz
# Repozytorium z zadaniami e-biznes
## Zadanie 1 - docker
[link do obrazu](https://hub.docker.com/r/jakubduj/zadanie_docker/tags)

## Zadanie 2 Scala

✅ 3.0 Należy stworzyć kontroler do Produktów [commit](https://github.com/Jakub-Danielewicz/ebiznes/tree/c55edaa407432e1bfa9b6c3d1dc0151c3530b2b0)

✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy [commit](https://github.com/Jakub-Danielewicz/ebiznes/tree/c55edaa407432e1bfa9b6c3d1dc0151c3530b2b0)

❌ 4.0 

❌ 4.5 

❌ 5.0 

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_scala)

## Zadanie 3 Kotlin

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord [commit](https://github.com/Jakub-Danielewicz/ebiznes/commit/071ecefd69e125113b1914dbe1e606c481e64b97)

✅ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota) [commit](https://github.com/Jakub-Danielewicz/ebiznes/commit/071ecefd69e125113b1914dbe1e606c481e64b97)

✅ 4.0 Zwróci listę kategorii na określone żądanie użytkownika [commit](https://github.com/Jakub-Danielewicz/ebiznes/commit/071ecefd69e125113b1914dbe1e606c481e64b97)

✅ 4.5 Zwróci listę produktów wg żądanej kategorii [commit](https://github.com/Jakub-Danielewicz/ebiznes/commit/071ecefd69e125113b1914dbe1e606c481e64b97)

❌ 5.0 

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_kotlin)

## Zadanie 4 Go

✅ 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD

✅ 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)

✅ 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

✅ 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem

❌/✅ 5.0 pogrupować zapytania w gorm’owe scope'y (tylko dla kategorii produktu)

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_go)


## Zadanie 5 Frontend

✅ 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej;

✅ 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing

✅ 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks

❌ 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose

✅ 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_frontend)

## Zadanie 6 Testy

✅ 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala)

✅ 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji

✅ 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami

✅ 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint

❌ 5.0 Należy uruchomić testy funkcjonalne na Browserstacku

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_cypress)

## Zadanie 7 Sonar

✅ 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
hookach gita [skrypt w backend/scripts]([https://github.com/Jakub-Danielewicz/ebiznes/commit/071ecefd69e125113b1914dbe1e606c481e64b97](https://github.com/Jakub-Danielewicz/ebiznes/blob/master/zadanie_cypress/scripts/golanglinter.sh))

✅ 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
aplikacji serwerowej)  [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=danielewicz_jakub_danielewicz&metric=bugs)](https://sonarcloud.io/summary/new_code?id=danielewicz_jakub_danielewicz)

✅ 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
aplikacji serwerowej) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=danielewicz_jakub_danielewicz&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=danielewicz_jakub_danielewicz)

✅ 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
w kodzie w Sonarze (kod aplikacji serwerowej) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=danielewicz_jakub_danielewicz&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=danielewicz_jakub_danielewicz) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=danielewicz_jakub_danielewicz&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=danielewicz_jakub_danielewicz)

✅ 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie 
aplikacji klienckiej  [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Jakub_Danielewicz&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Jakub_Danielewicz) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Jakub_Danielewicz&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Jakub_Danielewicz) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Jakub_Danielewicz&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Jakub_Danielewicz)

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_cypress)

## Zadanie 8 Oauth2

✅ 3.0 logowanie przez aplikację serwerową (bez Oauth2)

✅ 3.5 rejestracja przez aplikację serwerową (bez Oauth2)

✅ 4.0 logowanie via Google OAuth2

✅ 4.5 logowanie via Facebook lub Github OAuth2

✅ 5.0 zapisywanie danych logowania OAuth2 po stronie serwera

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_oauth2)

## Zadanie 9 ChatGPT bot

✅ 3.0 należy stworzyć po stronie serwerowej osobny serwis do łącznia z
chatGPT do usługi chat (wykorzystano google gemini)

✅ 3.5 należy stworzyć interfejs frontowy dla użytkownika, który
komunikuje się z serwisem; odpowiedzi powinny być wysyałen do
frontendowego interfejsu

❌/✅ 4.0 stworzyć listę 5 różnych otwarć oraz zamknięć rozmowy

✅ 4.5 filtrowanie po zagadnieniach związanych ze sklepem (np.
ograniczenie się jedynie do ubrań oraz samego sklepu) do GPT

❌ 5.0 filtrowanie odpowiedzi po sentymencie

Kod: [folder](https://github.com/Jakub-Danielewicz/ebiznes/tree/master/zadanie_LLM)

