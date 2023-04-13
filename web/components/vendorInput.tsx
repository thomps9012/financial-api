export default function VendorInput({
  state,
  setState,
}: {
  state: any;
  setState: any;
}) {
  const handleChange = (event: any) => {
    const { name, value } = event.target;
    const new_state = { ...state, [name]: value.trim().toLowerCase() };
    setState(new_state);
  };
  return (
    <form>
      <h2>Vendor Info</h2>
      <h3>Name</h3>
      <input
        defaultValue={state.name}
        type="text"
        name="name"
        onChange={handleChange}
      />
      <h4>Website</h4>
      <input
        defaultValue={state.website}
        type="text"
        name="website"
        onChange={handleChange}
      />
      <h4>Address</h4>
      <input
        defaultValue={state.address_line_one}
        type="text"
        name="address_line_one"
        onChange={handleChange}
      />
      <h4>Address Continued</h4>
      <input
        defaultValue={state.address_line_two}
        type="text"
        name="address_line_two"
        onChange={handleChange}
      />
      <br />
    </form>
  );
}
