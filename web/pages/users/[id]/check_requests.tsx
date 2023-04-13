import { GetServerSidePropsContext } from "next";

function UserCheckPage({ user_id }: { user_id: string }) {
  return (
    <main>
      <h1>Check page for {user_id}</h1>
    </main>
  );
}

UserCheckPage.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    user_id: id,
  };
};

export default UserCheckPage;
