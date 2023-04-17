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
    for (const field of Object.keys(checkRequestInfo)) {
      const input = document.getElementById(field) as
        | HTMLInputElement
        | HTMLSelectElement;
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
    }
    for (const field of Object.keys(vendor)) {
      const input = document.getElementById(
        "vendor-" + field
      ) as HTMLInputElement;
      if (
        input?.value != "" &&
        input?.value != null &&
        input?.value != undefined
      ) {
        document
          .getElementById(`invalid-vendor-${field}`)
          ?.classList.add("hidden");
        document
          .getElementById("vendor-" + field)
          ?.classList.remove("invalid-input");
      } else {
        document
          .getElementById("vendor-" + field)
          ?.classList.add("invalid-input");
        document
          .getElementById(`invalid-vendor-${field}`)
          ?.classList.remove("hidden");
      }
    }

    !new_request && request_id && fetchRequestInfo(request_id);
  }, [new_request, request_id, user_credentials, vendor, checkRequestInfo]);
  const validatePurchases = () => {
    const purchaseRows = document.getElementsByClassName("purchase-row");
    for (let i = 0; i < purchaseRows.length; i++) {
      const elements = purchaseRows[i].children;
      if ((elements[1] as HTMLInputElement).value === "") {
        elements[1].classList.add("invalid-input");
        elements[2].classList.remove("hidden");
      } else {
        elements[1].classList.remove("invalid-input");
        elements[2].classList.add("hidden");
      }
      if ((elements[4] as HTMLInputElement).value === "") {
        elements[4].classList.add("invalid-input");
        elements[5].classList.remove("hidden");
      } else {
        elements[4].classList.remove("invalid-input");
        elements[5].classList.add("hidden");
      }
      if (
        parseInt((elements[7] as HTMLInputElement).value) === 0 ||
        (elements[7] as HTMLInputElement).value === "" ||
        (elements[7] as HTMLInputElement).value === undefined
      ) {
        elements[7].classList.add("invalid-input");
        elements[8].classList.remove("hidden");
      } else {
        elements[7].classList.remove("invalid-input");
        elements[8].classList.add("hidden");
      }
    }
  };
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
    <form id="check-request-form">
      <CategorySelect state={checkRequestInfo} setState={setCheckRequestInfo} />
      <span id="invalid-category" className="REJECTED field-span">
        <br />
        Category is Required
      </span>
      <GrantSelect state={checkRequestInfo} setState={setCheckRequestInfo} />
      <span id="invalid-grant_id" className="REJECTED field-span">
        <br />
        Grant is Required
      </span>
      <h3>Purchase Date</h3>
      <input type="date" name="date" id="date" onChange={handleChange} />
      <span id="invalid-date" className="REJECTED field-span">
        <br />
        Purchase Date is Required
      </span>
      <h3>Description</h3>
      <textarea
        rows={5}
        maxLength={75}
        name="description"
        id="description"
        defaultValue={checkRequestInfo.description}
        onChange={handleChange}
      />
      <br />
      <span id="invalid-description" className="REJECTED field-span">
        Description is Required
      </span>
      <span>{checkRequestInfo.description.length}/75 characters</span>
      <VendorInput state={vendor} setState={setVendor} />
      <h2
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
        }}
      >
        Purchase List{" "}
        <span className="description-span">Limit 5 Purchases per Request</span>
      </h2>
      <PurchaseInput
        validatePurchase={validatePurchases}
        purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
      />
      {rowCount >= 2 && (
        <PurchaseInput
          validatePurchase={validatePurchases}
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 3 && (
        <PurchaseInput
          validatePurchase={validatePurchases}
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 4 && (
        <PurchaseInput
          validatePurchase={validatePurchases}
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      {rowCount >= 5 && (
        <PurchaseInput
          validatePurchase={validatePurchases}
          purchase={{ grant_line_item: "", description: "", amount: 0.0 }}
        />
      )}
      <div
        style={{
          display: "flex",
          justifyContent: "space-around",
          marginBottom: 10,
        }}
      >
        <a
          onClick={addPurchase}
          className="archive-btn"
          id="add_purchase"
          style={{ fontSize: 20 }}
        >
          + Purchase
        </a>
        <a
          onClick={removePurchase}
          className="reject-btn"
          id="remove_purchase"
          style={{ fontSize: 20 }}
        >
          - Purchase
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
        <option value="N/A">None</option>
        <option value="0380">5/3 Card Ending in 0380</option>
        <option value="6366">5/3 Card Ending in 6366</option>
      </select>
      <h2
        style={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "space-between",
        }}
      >
        Receipts{" "}
        <span className="description-span">
          Limit 5 Receipts in PNG or JPEG Format
        </span>
      </h2>
      <ReceiptUpload receipts={receipts} setReceipts={setReceipts} />
      <br />
      <a className="archive-btn" onClick={submitRequest}>
        Submit Request
      </a>
      <br />
    </form>
  );
}
