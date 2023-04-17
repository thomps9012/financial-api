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
    <span className="vendor-input">
      <h2>Vendor Info</h2>
      <h3>Name</h3>
      <input
        defaultValue={state.name}
        type="text"
        name="name"
        className="invalid-input"
        id="vendor-name"
        onChange={handleChange}
      />
      <span id="invalid-vendor-name" className="REJECTED field-span">
        <br />
        Vendor Name is Required
      </span>
      <h4>Address</h4>
      <input
        defaultValue={state.address_line_one}
        type="text"
        id="vendor-address_line_one"
        className="invalid-input"
        name="address_line_one"
        onChange={handleChange}
      />
      <span
        id="invalid-vendor-address_line_one"
        className="REJECTED field-span"
      >
        <br />
        Vendor Address is Required
      </span>
      <h4>Address Continued</h4>
      <input
        defaultValue={state.address_line_two}
        type="text"
        name="address_line_two"
        onChange={handleChange}
      />
      <h4>Website</h4>
      <input
        defaultValue={state.website}
        type="text"
        name="website"
        onChange={handleChange}
      />
      <br />
    </span>
  );
}
