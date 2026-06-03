export const useVitalSignService = () => {
  const { $api } = useNuxtApp();

  const fetchVitals = () => {
    return $api("/get-vitals", {
      method: "GET",
    });
  };

  return { fetchVitals };
};
