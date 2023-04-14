import { useAppContext } from "@/context/AppContext";
import { User_Name_Info } from "@/types/users";
import UnAuthorized from "./unAuthorized";
import { useEffect, useState } from "react";
import axios from "axios";

export default function UserSelect({ reportType }: { reportType: string }) {
  const [user_list, setUserList] = useState(new Array<User_Name_Info>());
  const { user_credentials, user_profile } = useAppContext();
  useEffect(() => {
    async function fetchUsers() {
      const { data } = await axios.get("/api/users", {
        ...user_credentials,
      });
      const user_data = data.data;
      setUserList(user_data);
    }
    user_profile.admin && fetchUsers();
  }, [user_profile, user_credentials]);
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
