import { useEffect, useState } from "react";
import axios from "axios";
import CategorySelect from "./categorySelect";
import GrantSelect from "./grantSelect";
import { useAppContext } from "@/context/AppContext";
import { Axios_Credentials } from "@/types/users";
export default function MileageForm({
  new_request,
  request_id,
}: {
  new_request: boolean;
  request_id?: string;
}) {
  const { user_credentials } = useAppContext();
  const [mileageRequestInput, setMileageRequestInput] = useState({
    grant_id: "",
    date: new Date().toISOString(),
    category: "",
    starting_location: "",
    destination: "",
    trip_purpose: "",
    start_odometer: 0,
    end_odometer: 1,
    tolls: 0.0,
    parking: 0.0,
  });
  useEffect(() => {
    const fetchRequestInfo = async (request_id: string) => {
      const { data } = await axios.get("/mileage/detail", {
        ...user_credentials,
        data: {
          mileage_id: request_id,
        },
      });
      const {
        grant_id,
        date,
        category,
        starting_location,
        destination,
        trip_purpose,
        start_odometer,
        end_odometer,
        tolls,
        parking,
      } = data.data;
      setMileageRequestInput({
        grant_id,
        date,
        category,
        starting_location,
        destination,
        trip_purpose,
        start_odometer,
        end_odometer,
        tolls,
        parking,
      });
    };
    !new_request && request_id && fetchRequestInfo(request_id);
  }, [new_request, request_id, user_credentials]);

  const handleChange = (e: any) => {
    e.preventDefault();
    const { name, value } = e.target;
    let new_state;
    switch (name.trim().toLowerCase()) {
      case "date":
        new_state = {
          ...mileageRequestInput,
          [name]: new Date(value).toISOString(),
        };
        break;
      case "start_odometer":
      case "end_odometer":
        new_state = {
          ...mileageRequestInput,
          [name]: parseInt(value),
        };
        break;
      case "tolls":
      case "parking":
        new_state = {
          ...mileageRequestInput,
          [name]: parseFloat(value),
        };
        break;
      default:
        new_state = {
          ...mileageRequestInput,
          [name]: value.trim().toLowerCase(),
        };
        break;
    }
    setMileageRequestInput(new_state);
  };
  const createMileage = async (config: Axios_Credentials) => {
    axios
      .post("/api/mileage", mileageRequestInput, config)
      .then((response) => console.log(response.data))
      .catch((error) => console.error(error));
  };
  const saveEdits = async (config: Axios_Credentials) => {
    const request_body = { ...mileageRequestInput, request_id };
    axios
      .put("/api/mileage", request_body, config)
      .then((res) => console.log(res))
      .catch((err) => console.error(err));
  };
  const handleSubmit = async (e: any) => {
    e.preventDefault();
    let res;
    if (new_request) {
      res = await createMileage(user_credentials);
    } else {
      res = await saveEdits(user_credentials);
    }
    console.log(res);
  };

  return (
    <form id="mileage-form">
      <GrantSelect
        state={mileageRequestInput}
        setState={setMileageRequestInput}
      />
      <CategorySelect
        state={mileageRequestInput}
        setState={setMileageRequestInput}
      />
      <h4>Trip Date</h4>
      <input type="datetime-local" name="date" onChange={handleChange} />
      <h4>Starting Location</h4>
      <input
        name="starting_location"
        defaultValue={mileageRequestInput.starting_location}
        id="start"
        maxLength={50}
        type="text"
        onChange={handleChange}
      />
      <br />
      <span>{mileageRequestInput.starting_location.length}/50 characters</span>
      <h4>Destination</h4>
      <input
        name="destination"
        id="end"
        defaultValue={mileageRequestInput.destination}
        maxLength={50}
        type="text"
        onChange={handleChange}
      />
      <span>{mileageRequestInput.destination.length}/50 characters</span>
      <h4>Trip Purpose</h4>
      <textarea
        rows={5}
        maxLength={75}
        name="trip_purpose"
        defaultValue={mileageRequestInput.trip_purpose}
        onChange={handleChange}
      />
      <span>{mileageRequestInput.trip_purpose.length}/75 characters</span>
      <br />
      <h4>Start Odometer</h4>
      <input
        name="start_odometer"
        defaultValue={mileageRequestInput.start_odometer}
        max={mileageRequestInput.end_odometer - 1}
        type="number"
        onChange={handleChange}
      />
      <h4>End Odometer</h4>
      <input
        name="end_odometer"
        defaultValue={mileageRequestInput.end_odometer}
        type="number"
        min={mileageRequestInput.start_odometer + 1}
        onChange={handleChange}
      />
      <h4>Tolls</h4>
      <input
        name="tolls"
        type="number"
        defaultValue={mileageRequestInput.tolls}
        onChange={handleChange}
      />
      <h4>Parking</h4>
      <input
        name="parking"
        type="number"
        defaultValue={mileageRequestInput.parking}
        onChange={handleChange}
      />
      <br />
      <a onClick={handleSubmit} className="archive-btn">
        Submit Request
      </a>
      <br />
    </form>
  );
}
