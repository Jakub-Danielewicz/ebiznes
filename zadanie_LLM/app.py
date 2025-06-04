from flask import Flask, render_template, request, redirect, url_for
from llm_service import ask_gemini
import os
from dotenv import load_dotenv
import google.generativeai as genai

app = Flask(__name__)

chat_history = []

@app.route("/", methods=["GET", "POST"])
def chat():
    if request.method == "POST":
        user_message = request.form["message"]
        bot_reply = ask_gemini(user_message, chat_history)
        chat_history.append(("Ty", user_message))
        chat_history.append(("Bot", bot_reply))
        return redirect(url_for("chat"))
    return render_template("chat.html", chat_history=chat_history)

if __name__ == "__main__":
    load_dotenv()
    genai.configure(api_key=os.environ.get("GEMINI_API_KEY"))
    for model in genai.list_models():
        print(model)
    app.run(debug=True)