import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { User_Name_Info } from "@/types/users";
import Link from "next/link";
import { useEffect, useState } from "react";
import axios from "axios";
export default function UserManagement() {
  const { user_profile } = useAppContext();
  const [user_list, setUserList] = useState(new Array<User_Name_Info>());
  const [filteredUsers, setFilteredUsers] = useState(user_list);
  useEffect(() => {
    user_profile.admin && fetchUsers();
  }, []);

  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  async function fetchUsers() {
    const { data } = await axios.get("/api/users");
    const user_data = data.data;
    setUserList(user_data);
  }
  const showHide = (e: any) => {
    e.preventDefault();
    const { id } = e.target;
    switch (id.trim().toLowerCase()) {
      case "all":
        setFilteredUsers(user_list);
        break;
      case "active":
        setFilteredUsers(user_list.filter(({ active }) => active));
        break;
      case "inactive":
        setFilteredUsers(user_list.filter(({ active }) => !active));
        break;
    }
  };
  return (
    <main>
      <h1>User List</h1>
      <section onClick={showHide}>
        <p>Show</p>
        <a className="archive-btn" id="all">
          All
        </a>
        <a className="archive-btn" id="active">
          ✅ Active
        </a>
        <a className="archive-btn" id="inactive">
          💤 Inactive
        </a>
      </section>
      {filteredUsers.map(({ id, name, active }) => (
        <article key={id}>
          <Link href={`/users/${id}`}>
            <h1 id={active ? `active` : `inactive`}>{name} Overview</h1>
          </Link>
          <Link href={`/users/${id}/edit`}>
            <h1 id={active ? `active` : `inactive`}>Edit {name} ✏️</h1>
          </Link>
          <Link href={`/users/${id}/mileage`}>
            <h1 id={active ? `active` : `inactive`}>
              🚗 {name} Mileage Requests
            </h1>
          </Link>
          <Link href={`/users/${id}/check`}>
            <h1 id={active ? `active` : `inactive`}>
              🗃️ {name} Check Requests
            </h1>
          </Link>
          <Link href={`/users/${id}/petty_cash`}>
            <h1 id={active ? `active` : `inactive`}>
              💵 {name} Petty Cash Requests
            </h1>
          </Link>
        </article>
      ))}
    </main>
  );
}
