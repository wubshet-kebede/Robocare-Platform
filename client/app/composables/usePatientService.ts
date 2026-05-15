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
  const admitPatient = (patientId: string) => {
    return $api("/admit-patient", {
      method: "POST",
      body: { patient_id: patientId },
    });
  };
  return { registerPatient, fetchPatients, admitPatient };
};
