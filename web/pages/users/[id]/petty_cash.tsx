import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function UserPettyCashPage({
  user_id,
  user_name,
  requests,
}: {
  user_id: string;
  user_name: string;
  requests: Petty_Cash_Overview[];
}) {
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>PettyCash for {user_name}</h1>
      <p>{JSON.stringify(requests, null, 2)}</p>
      <Link href={`/${user_id}`}>
        &larr; · &#8594; · &#x2192; Back to {user_name} Overview
      </Link>
    </main>
  );
}

UserPettyCashPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  const auth = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get("/api/user/petty_cash", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...auth,
      data: {
        error: data,
        user_id: auth.headers.Authorization.split(" ")[1],
        error_path: "GET /user/petty_cash",
        error_message: statusText,
      },
    });
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  const {
    data: user_data,
    status: user_status,
    statusText: user_statusText,
  } = await axios.get("/api/user/name", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  if (user_status != 200 || 201) {
    await axios.post("/api/error", {
      ...auth,
      data: {
        error: user_data,
        user_id: auth.headers.Authorization.split(" ")[1],
        error_path: "GET /user/petty_cash",
        error_message: user_statusText,
      },
    });
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  return {
    user_id: id,
    user_name: user_data.data,
    requests: data.data,
  };
};

export default UserPettyCashPage;
