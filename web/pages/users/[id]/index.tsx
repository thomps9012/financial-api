import { User_Public_Info } from "@/types/users";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

// export async function getStaticPaths(ctx: GetServerSidePropsContext) {
//   const credentials = getCookie("auth_credentials", {
//     req: ctx.req,
//     res: ctx.res,
//   });
//   const user_credentials = JSON.parse(credentials as string);
//   const { data } = await axios.get("/api/users", {
//     ...user_credentials,
//   });
//   const paths = data.data.map(({ id }: { id: string }) => ({
//     params: { id },
//   }));
//   return { paths, fallback: false };
// }

// export async function getStaticProps(ctx: GetServerSidePropsContext) {
//   const { id } = ctx.query;
//   const credentials = getCookie("auth_credentials", {
//     req: ctx.req,
//     res: ctx.res,
//   });
//   if (!credentials || credentials === "") {
//     return {
//       public_info: {},
//     };
//   }
//   const auth = JSON.parse(credentials as string);
//   const { data } = await axios.get("/user/detail", {
//     ...auth,
//     data: {
//       user_id: id,
//     },
//   });
//   return {
//     public_info: data.data,
//   };
// }
function UserOverviewPage({ public_info }: { public_info: User_Public_Info }) {
  return (
    <main>
      <h1>Overview page for {public_info?.name}</h1>
      <p>{JSON.stringify(public_info, null, 2)}</p>
      <Link href={`/${public_info?.id}/mileage`}>Mileage Requests</Link>
      <Link href={`/${public_info?.id}/check_requests`}>Check Requests</Link>
      <Link href={`/${public_info?.id}/petty_cash`}>Petty Cash Requests</Link>
    </main>
  );
}

UserOverviewPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      public_info: {},
    };
  }
  const auth = JSON.parse(credentials as string);
  const { data } = await axios.get("/user/detail", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  return {
    public_info: data.data,
  };
};

export default UserOverviewPage;
