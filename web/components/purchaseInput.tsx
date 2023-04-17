import { Purchase } from "../types/check_requests";

export default function PurchaseInput({
  purchase,
  validatePurchase,
}: {
  purchase: Purchase;
  validatePurchase: any;
}) {
  const { grant_line_item, description, amount } = purchase;
  return (
    <section className="purchase-row">
      <label>Grant Line Item</label>
      <input
        onChange={validatePurchase}
        defaultValue={grant_line_item}
        type="text"
        className="invalid-input"
        name="grant_line_item"
      />
      <span id="invalid-purchase" className="REJECTED field-span">
        <br />
        Must Include a Valid Purchase Grant Line Item
      </span>
      <label>Description</label>
      <input
        onChange={validatePurchase}
        defaultValue={description}
        type="text"
        className="invalid-input"
        name="description"
      />
      <span id="invalid-purchase" className="REJECTED field-span">
        <br />
        Must Include a Valid Purchase Description
      </span>
      <label>Amount</label>
      <input
        onChange={validatePurchase}
        defaultValue={amount}
        type="number"
        className="invalid-input"
        name="amount"
      />
      <span id="invalid-purchase" className="REJECTED field-span">
        <br />
        Must Include a Valid Purchase Amount
      </span>
      <br />
    </section>
  );
}
