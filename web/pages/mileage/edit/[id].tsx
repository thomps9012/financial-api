import { GetServerSidePropsContext } from "next";

function EditMileageRequest({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Edit Mileage Request {request_id}</h1>
    </main>
  );
}
EditMileageRequest.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};

export default EditMileageRequest;
