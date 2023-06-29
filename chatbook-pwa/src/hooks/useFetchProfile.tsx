import { useEffect, useState } from 'react';
import UserService from '../services/UserService';

const useFetchProfile = () => {
    const [profile, setProfile] = useState(null);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchProfile = async () => {
            setError(null);
            const response = await UserService.getProfile()

            console.log("Resp: ",response.statusText)
            if (response.status === 200) {
                const data = response.data;
                setProfile(data);
            };
        }

        fetchProfile();
    }, []);

    return { profile, error };
};

export default useFetchProfile;
