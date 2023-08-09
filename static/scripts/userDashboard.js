var viewMode = document.getElementById("viewMode");
viewMode.addEventListener("change", () => {
  var mode = viewMode.value;
  var a = document.createElement("a");
  if (mode == "available") a.href = "/userDashboard";
  else a.href = `/userDashboard/${mode}`;
  a.click();
});