import { GetServerSidePropsContext } from "next";

function UserEditPage({ user_id }: { user_id: string }) {
  return (
    <main>
      <h1>Edit page for {user_id}</h1>
    </main>
  );
}
UserEditPage.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    user_id: id,
  };
};

export default UserEditPage;
