export const useAdmittedPatientService = () => {
  const { $api } = useNuxtApp();
  const getAssignedPatients = () => {
    return $api("/assigned-patients", {
      method: "GET",
    });
  };
  return { getAssignedPatients };
};
