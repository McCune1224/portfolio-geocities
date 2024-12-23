function checkKeyPress(e) {
  if (e.key === "j" || e.key === "J") {
    document.cookie = "player_direction=down";
  }
  if (e.key === "k" || e.key === "K") {
    document.cookie = "player_direction=up";
  }
}
document.addEventListener("keypress", checkKeyPress);
