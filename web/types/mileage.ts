import { Action } from "./actions";

export interface Mileage_Request {
  id: string;
  grant_id: string;
  user_id: string;
  date: string;
  category: string;
  starting_location: string;
  destination: string;
  trip_purpose: string;
  start_odometer: number;
  end_odometer: number;
  tolls: number;
  parking: number;
  trip_mileage: number;
  reimbursement: number;
  created_at: string;
  action_history: Action[];
  current_user: string;
  current_status: string;
  is_active: boolean;
}

export interface Mileage_Overview {
  id: string;
  user_id: string;
  date: string;
  reimbursement: number;
  current_user: string;
  current_status: string;
  is_active: boolean;
}

export interface Mileage_Input {
  grant_id: string;
  date: string;
  category: string;
  starting_location: string;
  destination: string;
  trip_purpose: string;
  start_odometer: number;
  end_odometer: number;
  tolls: number;
  parking: number;
}

export interface Edit_Mileage_Input {
  id: string;
  user_id: string;
  grant_id: string;
  date: string;
  category: string;
  starting_location: string;
  destination: string;
  trip_purpose: string;
  start_odometer: number;
  end_odometer: number;
  tolls: number;
  parking: number;
}
