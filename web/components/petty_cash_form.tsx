import { useEffect, useState } from "react";
import CategorySelect from "./categorySelect";
import GrantSelect from "./grantSelect";
import ReceiptUpload from "./receiptUpload";
import axios from "axios";
import { useAppContext } from "@/context/AppContext";
import ErrorDisplay from "./errorDisplay";

export default function PettyCashFrom({
  new_request,
  request_id,
}: {
  new_request: boolean;
  request_id?: string;
}) {
  const { user_credentials } = useAppContext();
  const [pettyCashInput, setPettyCashInput] = useState({
    grant_id: "",
    category: "",
    date: new Date().toISOString(),
    description: "",
    amount: 0.0,
  });
  useEffect(() => {
    const fetchRequestInfo = async (request_id: string) => {
      const { data, status, statusText } = await axios.get(
        "/api/petty_cash/detail",
        {
          ...user_credentials,
          data: {
            petty_cash_id: request_id,
          },
        }
      );
      if (status != 200 || 201) {
        return (
          <ErrorDisplay
            message={statusText}
            path="GET /petty_cash/detail"
            error={data}
          />
        );
      }
      const { grant_id, category, date, description, amount, receipts } =
        data.data;
      setPettyCashInput({
        grant_id,
        category,
        date,
        description,
        amount,
      });
      setReceipts(receipts);
    };
    for (const field of Object.keys(pettyCashInput)) {
      const input = document.getElementById(field) as
        | HTMLSelectElement
        | HTMLInputElement;
      if (
        input?.value != "" &&
        input?.value != null &&
        input?.value != undefined
      ) {
        document.getElementById(`invalid-${field}`)?.classList.add("hidden");
        document.getElementById(field)?.classList.remove("invalid-input");
      } else {
        document.getElementById(field)?.classList.add("invalid-input");
        document.getElementById(`invalid-${field}`)?.classList.remove("hidden");
      }
      if (field === "amount") {
        if (
          parseInt(input.value) === 0 ||
          input.value === undefined ||
          input.value === ""
        ) {
          document.getElementById(field)?.classList.add("invalid-input");
          document
            .getElementById(`invalid-${field}`)
            ?.classList.remove("hidden");
        } else {
          document.getElementById(`invalid-${field}`)?.classList.add("hidden");
          document.getElementById(field)?.classList.remove("invalid-input");
        }
      }
    }
    !new_request && request_id && fetchRequestInfo(request_id);
  }, [new_request, request_id, user_credentials, pettyCashInput]);
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
      <CategorySelect state={pettyCashInput} setState={setPettyCashInput} />
      <span id="invalid-category" className="REJECTED field-span">
        <br />
        Category is Required
      </span>
      <GrantSelect state={pettyCashInput} setState={setPettyCashInput} />
      <span id="invalid-grant_id" className="REJECTED field-span">
        <br />
        Grant is Required
      </span>
      <h4>Request Amount</h4>
      <input
        type="number"
        id="amount"
        className="invalid-input"
        name="amount"
        defaultValue={pettyCashInput.amount}
        onChange={handleChange}
      />
      <span id="invalid-amount" className="REJECTED field-span">
        <br />A Valid Amount is Required
      </span>
      <h4>Date</h4>
      <input
        type="date"
        name="date"
        id="date"
        className="invalid-input"
        onChange={handleChange}
      />
      <span id="invalid-date" className="REJECTED field-span">
        <br />
        Request Date is Required
      </span>
      <h4>Description</h4>
      <textarea
        rows={5}
        maxLength={75}
        name="description"
        id="description"
        className="invalid-input"
        defaultValue={pettyCashInput.description}
        onChange={handleChange}
      />
      <br />
      <span id="invalid-description" className="REJECTED field-span">
        Request Description is Required
      </span>
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
