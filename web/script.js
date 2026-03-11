const list = document.getElementById("list");
const addItemBtn = document.getElementById("add-item");
const doneBtn = document.getElementById("done");
const editBtn = document.getElementById("edit");
const grandTotalEl = document.getElementById("grand-total");
const searchInput = document.getElementById("search");
const suggestionButtons = Array.from(document.querySelectorAll("[data-item]"));

const currency = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  minimumFractionDigits: 2,
});

let isCompleted = false;

function toNumber(value) {
  const num = parseFloat(value);
  return Number.isFinite(num) ? num : 0;
}

function formatCurrency(value) {
  return currency.format(value);
}

function updateRowTotal(row) {
  const qtyInput = row.querySelector(".qty-input");
  const priceInput = row.querySelector(".price-input");
  const totalCell = row.querySelector(".total-cell");
  const qty = toNumber(qtyInput.value);
  const price = toNumber(priceInput.value);
  const total = qty * price;
  totalCell.textContent = formatCurrency(total);
  updateGrandTotal();
}

function updateGrandTotal() {
  const totals = Array.from(list.querySelectorAll(".item-row"));
  const sum = totals.reduce((acc, row) => {
    const totalText = row.querySelector(".total-cell").textContent;
    const value = toNumber(totalText.replace(/[^0-9.]/g, ""));
    return acc + value;
  }, 0);
  grandTotalEl.textContent = formatCurrency(sum);
}

function applyReadOnlyState() {
  list.classList.toggle("completed", isCompleted);
  const inputs = list.querySelectorAll("input");
  inputs.forEach((input) => {
    input.readOnly = isCompleted;
  });
  searchInput.readOnly = isCompleted;
  addItemBtn.disabled = isCompleted;
  addItemBtn.classList.toggle("is-disabled", isCompleted);
  suggestionButtons.forEach((btn) => {
    btn.disabled = isCompleted;
  });
  doneBtn.hidden = isCompleted;
  editBtn.hidden = !isCompleted;
}

function createRow(name = "", qty = 1, price = 0) {
  const row = document.createElement("div");
  row.className = "item-row";

  row.innerHTML = `
    <input class="name-input" type="text" placeholder="Item name" value="${name}" />
    <input class="qty-input" type="number" min="0" step="1" value="${qty}" />
    <input class="price-input" type="number" min="0" step="0.01" value="${price}" />
    <div class="total-cell">${formatCurrency(qty * price)}</div>
  `;

  const qtyInput = row.querySelector(".qty-input");
  const priceInput = row.querySelector(".price-input");

  qtyInput.addEventListener("input", () => updateRowTotal(row));
  priceInput.addEventListener("input", () => updateRowTotal(row));

  const nameInput = row.querySelector(".name-input");
  nameInput.addEventListener("input", applySearchFilter);

  list.appendChild(row);
  applyReadOnlyState();
  updateGrandTotal();
}

function applySearchFilter() {
  const query = searchInput.value.trim().toLowerCase();
  const rows = list.querySelectorAll(".item-row");
  rows.forEach((row) => {
    const name = row.querySelector(".name-input").value.toLowerCase();
    row.style.display = name.includes(query) ? "grid" : "none";
  });
}

addItemBtn.addEventListener("click", () => {
  createRow("", 1, 0);
});

suggestionButtons.forEach((btn) => {
  btn.addEventListener("click", () => {
    const name = btn.dataset.item || "";
    createRow(name, 1, 0);
  });
});

searchInput.addEventListener("input", applySearchFilter);

doneBtn.addEventListener("click", () => {
  isCompleted = true;
  applyReadOnlyState();
});

editBtn.addEventListener("click", () => {
  isCompleted = false;
  applyReadOnlyState();
});

createRow("Milk", 1, 2.5);
createRow("Olive Oil", 1, 10.0);
createRow("Fresh Greens", 2, 3.5);
