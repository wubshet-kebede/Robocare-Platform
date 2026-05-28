export const useRoomService = () => {
  const { $api } = useNuxtApp();

  const registerRoom = (data: Record<string, any>) => {
    return $api("/register-room", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };
  const getRooms = () => {
    return $api("/get-rooms", {
      method: "GET",
    });
  };

  return { registerRoom, getRooms };
};
