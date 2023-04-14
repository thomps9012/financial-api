import { useEffect, useState } from "react";
import CategorySelect from "./categorySelect";
import GrantSelect from "./grantSelect";
import ReceiptUpload from "./receiptUpload";
import VendorInput from "./vendorInput";
import PurchaseInput from "./purchaseInput";
import { useAppContext } from "@/context/AppContext";
import axios from "axios";
export default function Check_Request_Form({
  new_request,
  request_id,
}: {
  new_request: boolean;
  request_id?: string;
}) {
  const { user_credentials } = useAppContext();
  const [checkRequestInfo, setCheckRequestInfo] = useState({
    grant_id: "",
    date: new Date().toISOString(),
    category: "",
    description: "",
    credit_card: "",
  });
  const [receipts, setReceipts] = useState(new Array<String>());
  const [vendor, setVendor] = useState({
    name: "",
    website: "",
    address_line_one: "",
    address_line_two: "",
  });
  useEffect(() => {
    const fetchRequestInfo = async (request_id: string) => {
      const { data } = await axios.get("/check/detail", {
        ...user_credentials,
        data: {
          check_request_id: request_id,
        },
      });
      const {
        grant_id,
        vendor,
        date,
        category,
        description,
        credit_card,
        purchases,
      } = data.data;
      setVendor(vendor);
      setCheckRequestInfo({
        grant_id,
        date,
        category,
        description,
        credit_card,
      });
      setRows(purchases.length);
      const purchase_inputs = document.getElementsByClassName("purchase-row");
      let i = 0;
      for (const purchase of purchases) {
        const purchase_input_fields = purchase_inputs[i];
        const { grant_line_item, description, amount } = purchase;
        purchase_input_fields.children[1].setAttribute(
          "value",
          grant_line_item
        );
        purchase_input_fields.children[3].setAttribute("value", description);
        purchase_input_fields.children[5].setAttribute("value", amount);
        i++;
      }
    };
    !new_request && request_id && fetchRequestInfo(request_id);
  }, [new_request, request_id, user_credentials]);

  const handleChange = (e: any) => {
    e.preventDefault();
    const { name, value } = e.target;
    let new_state;
    switch (name.trim().toLowerCase()) {
      case "date":
        new_state = {
          ...checkRequestInfo,
          [name]: new Date(value).toISOString(),
        };
        break;
      default:
        new_state = { ...checkRequestInfo, [name]: value.trim().toLowerCase() };
        break;
    }
    setCheckRequestInfo(new_state);
  };
  const [rowCount, setRows] = useState(1);
  const addPurchase = (e: any) => {
    e.preventDefault();
    rowCount < 5 ? setRows(rowCount + 1) : null;
  };
  const removePurchase = (e: any) => {
    e.preventDefault();
    setRows(rowCount - 1);
  };
  const createCheckRequest = async () => {};
  const saveEdits = async () => {
    const request_body = { ...checkRequestInfo, request_id };
  };
  const submitRequest = async (e: any) => {
    e.preventDefault();
    if (new_request) {
      const res = await createCheckRequest();
    } else {
      const res = await saveEdits();
    }
  };
  return (
    <form>
      <GrantSelect state={checkRequestInfo} setState={setCheckRequestInfo} />
      <CategorySelect state={checkRequestInfo} setState={setCheckRequestInfo} />
      <h3>Date</h3>
      <input type="date" name="date" onChange={handleChange} />
      <h3>Description</h3>
      <textarea
        rows={5}
        maxLength={75}
        name="description"
        defaultValue={checkRequestInfo.description}
        onChange={handleChange}
      />
      <span>{checkRequestInfo.description.length}/75 characters</span>
      <br />
      <VendorInput state={vendor} setState={setVendor} />
      <h2>Purchases</h2>
      <span className="description">Limit 5 Purchases per Request</span>
      <br />
      <PurchaseInput
        purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
      />
      {rowCount >= 2 && (
        <PurchaseInput
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 3 && (
        <PurchaseInput
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 4 && (
        <PurchaseInput
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 5 && (
        <PurchaseInput
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <a onClick={addPurchase} className="archive-btn" id="add_purchase">
          Add Purchase
        </a>
        <a onClick={removePurchase} className="reject-btn" id="remove_purchase">
          Remove Last
        </a>
      </div>
      <h3>Company Credit Card Used?</h3>
      <select
        name="credit_card"
        defaultValue={checkRequestInfo.credit_card}
        onChange={handleChange}
      >
        <option value="" disabled hidden>
          Select Credit Card..
        </option>
        <option value="N/A">No</option>
        <option value="1234">Card Ending in 1234</option>
        <option value="5678">Card Ending in 5678</option>
      </select>
      <h3>Receipts</h3>
      <span className="description">
        Limit 5 Receipts in PNG or JPEG Format
      </span>
      <ReceiptUpload receipts={receipts} setReceipts={setReceipts} />
      <br />
      <a className="archive-btn" onClick={submitRequest}>
        Submit Request
      </a>
      <br />
    </form>
  );
}
