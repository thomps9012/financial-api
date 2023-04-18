import { useAppContext } from "@/context/AppContext";
import { User_Name_Info } from "@/types/users";
import UnAuthorized from "./unAuthorized";
import { useEffect, useState } from "react";
import axios from "axios";
import ErrorDisplay from "./errorDisplay";
export default function UserSelect({
  reportType,
  setRequests,
}: {
  reportType: string;
  setRequests: any;
}) {
  const [user_list, setUserList] = useState(new Array<User_Name_Info>());
  const [selected_user, setUser] = useState("");
  const { user_credentials, user_profile } = useAppContext();
  useEffect(() => {
    async function fetchUsers() {
      const { data, status, statusText } = await axios.get("/api/users", {
        ...user_credentials,
      });
      if (status != 200 || 201) {
        return (
          <ErrorDisplay path="GET /users" message={statusText} error={data} />
        );
      }
      const user_data = data.data;
      setUserList(user_data);
    }
    user_profile.admin && fetchUsers();
  }, [user_profile, user_credentials]);
  let handleSubmit = async (e: any) => {
    e.preventDefault();
    if (selected_user === "") {
      document.getElementById("invalid-user")?.classList.remove("hidden");
      return;
    }
    document.getElementById("invalid-user")?.classList.add("hidden");
    const { data, status, statusText } = await axios.get(`/api/user/${reportType}`, {
      ...user_credentials,
      data: {
        user_id: selected_user,
      },
    });
    if (status != 200 || 201) {
      return (
        <ErrorDisplay path={`GET /users/${reportType}`} message={statusText} error={data} />
      );
    }
    setRequests(data.data);
  };
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <form>
      <select defaultValue={""} onChange={(e: any) => setUser(e.target.value)}>
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
      <span id="invalid-user" className="REJECTED field-span hidden">
        <br />
        Must Select a User
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
