import { useState } from "react";
import CategorySelect from "./categorySelect";
import GrantSelect from "./grantSelect";
import ReceiptUpload from "./receiptUpload";

export default function PettyCashFrom({
  new_request,
  request_id,
}: {
  new_request: boolean;
  request_id?: string;
}) {
  const [pettyCashInput, setPettyCashInput] = useState({
    grant_id: "",
    category: "",
    date: "",
    description: "",
    amount: 0.0,
  });
  const [receipts, setReceipts] = useState(new Array<String>());
  const handleChange = (e: any) => {
    e.preventDefault();
    const { name, value } = e.target;
    let new_state;
    switch (name.trim().toLowerCase()) {
      case "amount":
        new_state = { ...pettyCashInput, [name]: parseFloat(value) };
        break;
      default:
        new_state = { ...pettyCashInput, [name]: value.trim().toLowerCase() };
        break;
    }
    setPettyCashInput(new_state);
  };
  const createPettyCashRequest = async () => {};
  const saveEdits = async () => {
    const request_body = { ...pettyCashInput, request_id };
  };
  const handleSubmit = async (e: any) => {
    e.preventDefault();
    if (new_request) {
      const res = await createPettyCashRequest();
    } else {
      const res = await saveEdits();
    }
  };
  return (
    <form id="petty-cash-form">
      <GrantSelect state={pettyCashInput} setState={setPettyCashInput} />
      <CategorySelect state={pettyCashInput} setState={setPettyCashInput} />
      <h4>Amount</h4>
      <input
        type="number"
        name="amount"
        defaultValue={pettyCashInput.amount}
        onChange={handleChange}
      />
      <h4>Date</h4>
      <input
        type="date"
        name="date"
        defaultValue={pettyCashInput.date}
        onChange={handleChange}
      />
      <h4>Description</h4>
      <textarea
        rows={5}
        maxLength={75}
        name="description"
        defaultValue={pettyCashInput.description}
        onChange={handleChange}
      />
      <br />
      <span>{pettyCashInput.description.length}/75 characters</span>
      <br />
      <h3>Receipts</h3>
      <span className="description">
        Limit 5 Receipts in PNG or JPEG Format
      </span>
      <ReceiptUpload receipts={receipts} setReceipts={setReceipts} />
      <br />
      <a className="archive-btn" onClick={handleSubmit}>
        Submit Request
      </a>
      <br />
    </form>
  );
}
