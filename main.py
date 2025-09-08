from fastapi import FastAPI

app = FastAPI()

@app.get("/teams")
def read_teams():
    return {"teams": ["Real Madrid", "Barcelona", "Liverpool"]}

@app.get("/players")
def read_players():
    return {"players": ["Ronaldo", "Falcao", "Messi"]}
