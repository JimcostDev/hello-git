from fastapi.testclient import TestClient
from main import app

def test_read_teams():
    client = TestClient(app)
    response = client.get("/teams")
    assert response.status_code == 200
    data = response.json()
    assert "teams" in data
    assert "Real Madrid" in data["teams"]

def test_read_players():
    client = TestClient(app)
    response = client.get("/players")
    assert response.status_code == 200
    data = response.json()
    assert "players" in data
    assert "Messi" in data["players"]
