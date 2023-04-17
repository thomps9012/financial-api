import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function UserPettyCashPage({
  user_id,
  requests,
}: {
  user_id: string;
  requests: Petty_Cash_Overview[];
}) {
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>PettyCash page for {user_id}</h1>
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}

UserPettyCashPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  const auth = JSON.parse(credentials as string);
  const { data } = await axios.get("/user/petty_cash", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  if (credentials) {
    return {
      user_id: id,
      requests: data.data,
    };
  } else {
    return {
      user_id: "",
      requests: [],
    };
  }
};

export default UserPettyCashPage;
