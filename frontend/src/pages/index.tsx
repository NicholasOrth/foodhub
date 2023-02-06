
import Router from 'next/router'

import styles from '@/styles/Home.module.css'

export default function Home() {
  return (
      <div className={styles.pageContainer}>
          <div className={styles.authBtn}
               onClick={() => Router.push('/auth/login') }
          >
              <h3>Login</h3>
          </div>
          <div className={styles.authBtn}
               onClick={() => Router.push('/auth/signup') }
          >
              <h3>Sign Up</h3>
          </div>
      </div>
  )
}
