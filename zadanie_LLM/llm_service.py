import google.generativeai as genai
import os
import re

genai.configure(api_key=os.environ.get("GEMINI_API_KEY"))

SYSTEM_PROMPT = (
    "Jesteś asystentem sklepu z gitarami."
    "Odpowiadaj wyłącznie na pytania dotyczące sklepu oraz produktów w nim. Możesz używać swojej ogólnodostępnej wiedzy na temat produktów i kategorii, które są dostępne. "
    "Jeśli potrzebujesz informacji z bazy danych sklepu, nie generuj odpowiedzi na pytanie, tylko  napisz [call:getcategories] lub [call:getproducts:nazwa_kategorii]. "
    "Ceny produktów są podane po ich nazwie, np Gibson SG[2500zł], ale nie wyświetlaj cen, jeśli klient o nie nie zapyta."
    "Jeśli klient zapyta o cenę produktu możesz go znaleźć używając [call:getproducts:nazwa_kategorii]."
    "Spróbuj dopasować podaną przez klienta nazwę kategorii do nazwy kategorii zwróconej przez [call:getcategories]."
    "Po otrzymaniu odpowiedzi na wywołanie funkcji, dokończ odpowiedź na pytanie użytkownika. "
    "Jeśli pytanie nie dotyczy tych tematów, grzecznie odmów odpowiedzi. "
    "Odpowiadaj zawsze uprzejmie i grzecznie."
    "Swoją pierwszą odpowiedź w konwersacji zacznij od jednego z następujących otwarć: "
    "-Witamy w naszym sklepie, gdzie fale dźwiękowe ucieszą każdego!"
    "-Świetnie widzieć Cię w naszym sklepie, znajdziesz u nas najlepsze produkty gitarowe na całym świecie!"
    "-Wspaniale jest móc Cię zobaczyć w naszym sklepie! Na pewno znajdziesz tu produkty gitarowe dla siebie!"
    "-Witaj w sklepie gitarowym, gdzie Twoje brzmienie jest dla nas najważniejsze!"
    "-Witaj w sklepie! Cieszymy się, że postawiłeś na najlepsze brzmienie swoich instrumentów!"
    "Nie zaczynaj pozostałych odpowiedzi od otwarć."    
)

def getproducts(category="Gitary"):
    if category == "Gitary":
        return "Gibson SG[2500zł], Fender Stratocaster[4000zł]"
    if category == "Efekty Gitarowe":
        return "Boss GX-100[2000zł], Digitech Drop[1500zł], Digitech Whammy[1200zł], Dunlop CGB95[600zł], Strymon Mobius[700zł]"
    if category == "Wzmacniacze":
        return "Mesa Boogie Mark V[5000zł], Marshall JCM800[6000zł]"
    else:
        return "Podana kategoria jest pusta"

def getcategories():
    return "Gitary, Efekty Gitarowe, Wzmacniacze"

def ask_gemini(prompt, history=None):
    if history is None:
        history = []
    conversation = ""
    for who, msg in history:
        if who == "Ty":
            conversation += f"\tUżytkownik: {msg}\n"
        else:
            conversation += f"\tAsystent: {msg}\n"
    full_prompt = (
        f"{SYSTEM_PROMPT}\n\n"
        "historia konwersacji:\n"
        f"{conversation}\n"
        f"Nowe Pytanie: {prompt}\n"
        f"Asystent:"
    )
    print(full_prompt)
    model = genai.GenerativeModel("models/gemini-2.0-flash")
    response = model.generate_content(full_prompt)
    text = response.text.strip()

    call_match = re.search(r"\[call:(getcategories|getproducts)\s*(?::\s*([^\]\n]+))?\]", text, re.IGNORECASE | re.DOTALL)
    print("LLM odpowiedź:", text)
    print("Dopasowanie function call:", call_match)
    if call_match:
        func = call_match.group(1)
        arg = call_match.group(2)
        if func == "getcategories":
            result = getcategories()
            followup_prompt = (
                f"{SYSTEM_PROMPT}\n\n"
                "historia konwersacji:\n"
                f"{conversation}\n"
                f"Pytanie użytkownika: {prompt}\n"
                f"Odpowiedź funkcji getcategories: {result}\n"
                f"Dokończ odpowiedź dla użytkownika na podstawie powyższych informacji."
            )
        elif func == "getproducts" and arg:
            print("Wywołano funkcję getproducts z argumentem:", arg)
            result = getproducts(arg.strip())
            followup_prompt = (
                f"{SYSTEM_PROMPT}\n\n"
                "historia konwersacji:\n"
                f"{conversation}\n"
                f"Pytanie użytkownika: {prompt}\n"
                f"Odpowiedź funkcji getproducts({arg.strip()}): {result}\n"
                f"Dokończ odpowiedź dla użytkownika na podstawie powyższych informacji."
            )
        else:
            return "Nieprawidłowe wywołanie funkcji przez LLM."

        followup_response = model.generate_content(followup_prompt)
        return followup_response.text.strip()

    return text