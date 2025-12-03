import os
import google.generativeai as genai
from search import search


ANSWER_MODEL = "gemini-flash-latest"


def build_context(results):
    context = ""
    for score, r in results:
        context += f"\n---\n"
        context += f"Repo: {r['repo']}\n"
        context += f"File: {r['file']}\n"
        context += f"Code:\n{r.get('content', '')}\n"
    return context


def answer_question(question):
    results = search(question, top_k=3)

    if not results:
        return "No relevant context found."

    context = build_context(results)

    prompt = f"""
You are a senior software engineer.

Answer the question ONLY using the context below.
If the answer is not present, say: "I don’t have enough context."

Context:
{context}

Question:
{question}
"""

    model = genai.GenerativeModel(ANSWER_MODEL)
    response = model.generate_content(prompt)

    return response.text


def main():
    api_key = os.getenv("GEMINI_API_KEY")
    if not api_key:
        raise RuntimeError("GEMINI_API_KEY not set")

    genai.configure(api_key=api_key)

    question = input("Ask about the codebase: ")
    answer = answer_question(question)

    print("\nAnswer:\n")
    print(answer)


if __name__ == "__main__":
    main()
