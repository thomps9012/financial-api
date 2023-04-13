import { GetServerSidePropsContext } from "next";

function UserPettyCashPage({ user_id }: { user_id: string }) {
  return (
    <main>
      <h1>PettyCash page for {user_id}</h1>
    </main>
  );
}

UserPettyCashPage.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    user_id: id,
  };
};

export default UserPettyCashPage;
