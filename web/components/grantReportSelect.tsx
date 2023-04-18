import { useAppContext } from "@/context/AppContext";
import { Grant } from "@/types/grants";
import axios from "axios";
import { useState } from "react";
import ErrorDisplay from "./errorDisplay";

export default function GrantReportSelect({
  reportType,
  setReport,
}: {
  reportType: string;
  setReport: any;
}) {
  const { user_credentials } = useAppContext();
  const [grant_id, setGrantID] = useState("");
  const handleSubmit = async (e: any) => {
    e.preventDefault();
    if (grant_id === "") {
      document.getElementById("invalid-grant_id")?.classList.remove("hidden");
      return;
    }
    document.getElementById("invalid-grant_id")?.classList.add("hidden");
    const { data, status, statusText } = await axios.get(
      "/api/grant/" + reportType,
      {
        ...user_credentials,
        data: {
          id: grant_id,
        },
      }
    );
    if (status != 200 || 201) {
      return (
        <ErrorDisplay
          message={statusText}
          path={`GET /grant/${reportType}`}
          error={data}
        />
      );
    }
    setReport(data.data);
  };
  const grants = [
    {
      id: "H79TI082369",
      name: "BCORR",
    },
    {
      id: "SOR_HOUSING",
      name: "SOR Recovery Housing",
    },
    {
      id: "SOR_PEER",
      name: "SOR Peer",
    },
    {
      id: "SOR_LORAIN",
      name: "SOR Lorain 2.0",
    },
    {
      id: "H79TI085495",
      name: "RAP AID (Recover from Addition to Prevent Aids)",
    },
    {
      id: "2020-JY-FX-0014",
      name: "JSBT (OJJDP) - Jumpstart For A Better Tomorrow",
    },
    {
      id: "H79SP082264",
      name: "HIV Navigator",
    },
    {
      id: "H79SP082475",
      name: "SPF (HOPE 1)",
    },
    {
      id: "SOR_TWR",
      name: "SOR 2.0 - Together We Rise",
    },
    {
      id: "H79TI083370",
      name: "BSW (Bridge to Success Workforce)",
    },
    {
      id: "H79SM085150",
      name: "CCBHC",
    },
    {
      id: "H79TI083662",
      name: "IOP New Syrenity Intensive outpatient Program",
    },
    {
      id: "TANF",
      name: "TANF",
    },
    {
      id: "H79SP081048",
      name: "STOP Grant",
    },
    {
      id: "H79TI085410",
      name: "N MAT (NORA Medication-Assisted Treatment Program)",
    },
  ];
  return (
    <form>
      <select
        defaultValue={""}
        onChange={(e: any) => setGrantID(e.target.value)}
      >
        <option value="" disabled hidden>
          Select Grant...
        </option>
        <option value="N/A">None</option>
        {grants.map((grant: Grant) => {
          const { id, name } = grant;
          return (
            <option key={id} value={id}>
              {name}
            </option>
          );
        })}
      </select>
      <span id="invalid-grant_id" className="REJECTED field-span hidden">
        <br />
        Grant is Required
      </span>
      <br />
      <a
        className="archive-btn"
        onClick={handleSubmit}
        style={{ textAlign: "right" }}
      >
        Generate Report
      </a>
    </form>
  );
}
