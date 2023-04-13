import { useAppContext } from "@/context/AppContext";
import { Grant } from "@/types/grants";

export default function GrantReportSelect({
  reportType,
  setReport,
}: {
  reportType: string;
  setReport: any;
}) {
  const { grant_list } = useAppContext();
  const handleSubmit = (e: any) => {
    e.preventDefault();
  };
  return (
    <form>
      <h3>Grant {reportType} Report</h3>
      <select>
        <option value="" disabled hidden>
          Select Grant...
        </option>
        <option value="N/A">None</option>
        {grant_list.map((grant: Grant) => {
          const { id, name } = grant;
          return (
            <option key={id} value={id}>
              {name}
            </option>
          );
        })}
      </select>
      <a className="archive-btn" onClick={handleSubmit}>
        Generate Report
      </a>
    </form>
  );
}
