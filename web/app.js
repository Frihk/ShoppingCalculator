const state = {
  items: [],
};

const form = document.getElementById("item-form");
const nameInput = document.getElementById("item-name");
const qtyInput = document.getElementById("item-qty");
const priceInput = document.getElementById("item-price");
const clearBtn = document.getElementById("clear-items");
const checkoutBtn = document.getElementById("checkout");
const itemsList = document.getElementById("items-list");
const totalQty = document.getElementById("total-qty");
const totalCost = document.getElementById("total-cost");
const statusEl = document.getElementById("status");
const suggestionsList = document.getElementById("suggestions-list");

function renderItems() {
  if (state.items.length === 0) {
    itemsList.innerHTML = '<div class="list-empty">No items yet. Add your first item to begin.</div>';
    totalQty.textContent = "0";
    totalCost.textContent = "0.00";
    return;
  }

  let qty = 0;
  let cost = 0;

  itemsList.innerHTML = state.items
    .map((item, index) => {
      const lineTotal = item.quantity * item.price;
      qty += item.quantity;
      cost += lineTotal;

      return `
        <div class="list-row">
          <span>${item.name}</span>
          <span>${item.quantity}</span>
          <span>${item.price.toFixed(2)}</span>
          <span>${lineTotal.toFixed(2)}</span>
          <button type="button" data-remove="${index}">×</button>
        </div>
      `;
    })
    .join("");

  totalQty.textContent = String(qty);
  totalCost.textContent = cost.toFixed(2);
}

function setStatus(message, type = "") {
  statusEl.textContent = message;
  statusEl.classList.remove("success", "error");
  if (type) {
    statusEl.classList.add(type);
  }
}

async function loadSuggestions() {
  try {
    const response = await fetch("/api/suggestions");
    if (!response.ok) {
      throw new Error("Failed to load suggestions");
    }

    const data = await response.json();
    const suggestions = data.suggestions || [];

    if (suggestions.length === 0) {
      suggestionsList.innerHTML = '<div class="list-empty">No suggestions yet.</div>';
      return;
    }

    suggestionsList.innerHTML = suggestions
      .map(
        (item) => `
          <div class="suggestion-item">
            <span>${item.name}</span>
            <span>${Number(item.price).toFixed(2)} (${item.freq}x)</span>
          </div>
        `,
      )
      .join("");
  } catch (_err) {
    suggestionsList.innerHTML = '<div class="list-empty">Unable to load suggestions.</div>';
  }
}

form.addEventListener("submit", (event) => {
  event.preventDefault();

  const name = nameInput.value.trim();
  const quantity = Number(qtyInput.value);
  const price = Number(priceInput.value);

  if (!name || quantity <= 0 || price <= 0) {
    setStatus("Please provide valid item details.", "error");
    return;
  }

  state.items.push({ name, quantity, price });
  form.reset();
  qtyInput.value = "1";
  setStatus("Item added to list.", "success");
  renderItems();
});

itemsList.addEventListener("click", (event) => {
  const target = event.target;
  if (!(target instanceof HTMLElement)) {
    return;
  }

  const index = target.getAttribute("data-remove");
  if (index === null) {
    return;
  }

  state.items.splice(Number(index), 1);
  renderItems();
});

clearBtn.addEventListener("click", () => {
  state.items = [];
  renderItems();
  setStatus("List cleared.");
});

checkoutBtn.addEventListener("click", async () => {
  if (state.items.length === 0) {
    setStatus("Add at least one item before checkout.", "error");
    return;
  }

  try {
    const response = await fetch("/api/checkout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ items: state.items }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Checkout failed");
    }

    totalQty.textContent = String(data.totalQuantity);
    totalCost.textContent = Number(data.totalCost).toFixed(2);
    setStatus("Checkout saved successfully.", "success");

    state.items = [];
    renderItems();
    await loadSuggestions();
  } catch (err) {
    setStatus(err.message || "Checkout failed.", "error");
  }
});

renderItems();
loadSuggestions();
