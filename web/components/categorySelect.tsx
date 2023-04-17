export default function CategorySelect({
  state,
  setState,
  valid,
}: {
  valid: boolean;
  state: any;
  setState: any;
}) {
  const categories = [
    "IOP",
    "INTAKE",
    "PEERS",
    "ACT_TEAM",
    "IHBT",
    "PERKINS",
    "MENS_HOUSE",
    "NEXT_STEP",
    "LORAIN",
    "PREVENTION",
    "ADMINISTRATIVE",
    "FINANCE",
  ];
  const handleChange = (event: any) => {
    const { value } = event.target;
    const new_state = { ...state, category: value.trim().toUpperCase() };
    setState(new_state);
  };
  return (
    <>
      <h4>Category</h4>
      <select
        name={state}
        id="category"
        onChange={handleChange}
        defaultValue={state.category}
      >
        <option value="" disabled hidden>
          Select Category...
        </option>
        {categories.map((category) => (
          <option value={category} key={category}>
            {category.split("_").join(" ")}
          </option>
        ))}
      </select>
    </>
  );
}
