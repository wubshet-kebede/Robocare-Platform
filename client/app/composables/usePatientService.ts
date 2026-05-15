export const usePatientService = () => {
  const { $api } = useNuxtApp();

  const registerPatient = (data: Record<string, any>) => {
    return $api("/patients", {
      method: "POST",
      body: data,
    });
  };

  return { registerPatient };
};
