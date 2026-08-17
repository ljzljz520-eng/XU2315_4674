const status = document.querySelector("#status");
const safety = document.querySelector("#safety");
const featured = document.querySelector("#featured");
const team = document.querySelector("#team");
const form = document.querySelector("#search-form");

function articleCard(article) {
  const card = document.createElement("article");
  const section = document.createElement("p");
  section.className = "section";
  section.textContent = article.section;
  const title = document.createElement("h3");
  title.textContent = article.title;
  const summary = document.createElement("p");
  summary.textContent = article.summary;
  card.replaceChildren(section, title, summary);
  return card;
}

async function loadHome() {
  const response = await fetch("/api/home");
  const home = await response.json();
  safety.replaceChildren(...home.safetyTips.map((tip) => {
    const item = document.createElement("li");
    item.textContent = tip;
    return item;
  }));
  featured.replaceChildren(...home.featured.map(articleCard));
  team.replaceChildren(...home.authorTeam.map((author) => {
    const item = document.createElement("li");
    item.textContent = author;
    return item;
  }));
}

async function loadBasics() {
  if (location.pathname !== "/equipment-basics") {
    return;
  }
  const response = await fetch("/api/equipment-basics");
  const result = await response.json();
  const heading = document.createElement("h2");
  heading.textContent = "基础装备清单";
  const list = document.createElement("div");
  list.className = "article-grid";
  list.replaceChildren(...result.items.map((item) => {
    const card = document.createElement("article");
    const category = document.createElement("p");
    category.className = "section";
    category.textContent = item.category;
    const name = document.createElement("h3");
    name.textContent = item.name;
    const purpose = document.createElement("p");
    purpose.textContent = item.purpose;
    card.replaceChildren(category, name, purpose);
    return card;
  }));
  status.replaceChildren(heading, list);
}

form.addEventListener("submit", async (event) => {
  event.preventDefault();
  const query = new FormData(form).get("q");
  const response = await fetch(`/api/search?q=${encodeURIComponent(query)}`);
  const result = await response.json();
  if (result.articles.length === 0) {
    const message = document.createElement("p");
    message.append(`没有找到“${result.query}”的文章。可以先查看`);
    const link = document.createElement("a");
    link.href = result.starterListPath;
    link.textContent = "基础装备清单";
    message.append(link, "。");
    status.replaceChildren(message);
    return;
  }
  status.replaceChildren(...result.articles.map(articleCard));
});

Promise.all([loadHome(), loadBasics()]).catch(() => {
  status.textContent = "暂时无法加载笔记，请稍后再试。";
});
