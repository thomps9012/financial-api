import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { User_Name_Info } from "@/types/users";
import Link from "next/link";
import { useState } from "react";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function UserManagement({ user_list }: { user_list: User_Name_Info[] }) {
  const { user_profile } = useAppContext();
  const [filteredUsers, setFilteredUsers] = useState(user_list);
  if (!user_profile.admin) {
    return <UnAuthorized />;
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

UserManagement.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      user_list: [],
    };
  }
  const user_credentials = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get("/api/users", {
    ...user_credentials,
  });
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /users",
        error_message: statusText,
      },
    });
    return {
      user_list: [],
    };
  }
  return {
    user_list: data.data,
  };
};

export default UserManagement;
