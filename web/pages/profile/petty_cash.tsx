import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Request } from "@/types/petty_cash";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfilePettyCashPage({
  petty_cash_requests,
}: {
  petty_cash_requests: Petty_Cash_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Petty Cash Page for {name}</h1>
      <p>{JSON.stringify(petty_cash_requests, null, 2)}</p>
    </main>
  );
}

ProfilePettyCashPage.getInitialProps = async (
  ctx: GetServerSidePropsContext
) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  const user_credentials = JSON.parse(credentials as string);
  if (!credentials) {
    return {
      petty_cash_requests: [],
    };
  }
  const { data, status, statusText } = await axios.get(
    "/api/me/petty_cash",
    ...user_credentials
  );
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /me/petty_cash",
        error_message: statusText,
      },
    });
    return {
      petty_cash_requests: [],
    };
  }
  return {
    petty_cash_requests: data.data,
  };
};

export default ProfilePettyCashPage;
