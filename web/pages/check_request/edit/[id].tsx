import { GetServerSidePropsContext } from "next";
import styles from "../../../styles/Home.module.css";

function EditCheckRequest({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Check Request Edit Page for {request_id}</h1>
    </main>
  );
}

EditCheckRequest.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};

export default EditCheckRequest;
