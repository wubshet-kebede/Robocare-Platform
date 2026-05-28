export const usePatientService = () => {
  const { $api } = useNuxtApp();

  const registerPatient = (data: Record<string, any>) => {
    return $api("/register-patient", {
      method: "POST",
      body: data,
    });
  };
  const fetchPatients = () => {
    return $api("/fetch-patients", {
      method: "GET",
    });
  };
  const admitPatient = (data: Record<string, any>) => {
    return $api("/admit-patient", {
      method: "POST",
      body: data,
    });
  };
  return { registerPatient, fetchPatients, admitPatient };
};
