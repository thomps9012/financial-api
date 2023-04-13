import { useAppContext } from "@/context/AppContext";
import { User_Name_Info } from "@/types/users";
import UnAuthorized from "./unAuthorized";

export default function UserSelect({ reportType }: { reportType: string }) {
  let handleSubmit = (e: any) => {
    e.preventDefault();
    switch (reportType.trim().toLowerCase().split(" ").join("_")) {
      case "check":
        break;
      case "mileage":
        break;
      case "petty_cash":
        break;
      default:
        break;
    }
  };
  const { user_list, user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <form>
      <h3>User {reportType} Report</h3>
      <select>
        <option value="" disabled hidden>
          Select User...
        </option>
        {user_list.map((user: User_Name_Info) => {
          const { id, name } = user;
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
