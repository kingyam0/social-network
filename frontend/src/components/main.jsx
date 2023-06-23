import { Login } from './login';
import { Register } from './register';
import { Home } from './home';
import { LoggedIn } from './loggedIn/loggedIn'
import { Profile } from './profile/profile'
import { Posts } from './posts/posts'

export default function Main() {
    return (
        //  let [isLogged, setIsLogged] = useState(false);
        <>
            <Login />
            <Register />
            <Home />
            <LoggedIn />
            <Profile />
            <Posts />
        </>
    )
}