import { useAppContext } from "@/context/AppContext";
import axios from "axios";
import ErrorDisplay from "./errorDisplay";

export default function MonthlyReportSelect({
  reportType,
  setReport,
}: {
  reportType: string;
  setReport: any;
}) {
  const { user_credentials } = useAppContext();
  const handleSubmit = async (e: any) => {
    e.preventDefault();
    const date_input = (
      document.getElementById("month_select") as HTMLInputElement
    ).value.split("-");
    const month = parseInt(date_input[1]);
    const year = parseInt(date_input[0]);
    if (Number.isNaN(month) || Number.isNaN(year)) {
      document.getElementById("invalid-date")?.classList.remove("hidden");
      return;
    }
    document.getElementById("invalid-date")?.classList.add("hidden");
    const { data, status, statusText } = await axios.get(
      `/api/${reportType}/monthly`,
      {
        ...user_credentials,
        data: {
          month,
          year,
        },
      }
    );
    if (status != 200 || 201) {
      return (
        <ErrorDisplay
          message={statusText}
          path={`GET /${reportType}/monthly`}
          error={data}
        />
      );
    }
    setReport(data.data);
  };
  return (
    <form style={{ textAlign: "right" }}>
      <input type="month" id="month_select" name="month_select" />
      <label htmlFor="month_select">(month and year)</label>
      <span id="invalid-date" className="REJECTED field-span hidden">
        <br />
        Month & Year are Required
      </span>
      <br />
      <a onClick={handleSubmit} className="archive-btn">
        {" "}
        Generate Report{" "}
      </a>
    </form>
  );
}
