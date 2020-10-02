import React from 'react';

import './MsgBoard.css'

export default class MsgBoard extends React.Component {
    render() {
        return (
            <div>
                {this.props.msg === 'warning' &&
                    <div className="alert is-error">
                        You have exceeded the maximum number of users for your account, please upgrade your plan to increase the limit.
                    </div>
                }

                {this.props.msg === 'success' &&
                    <div className="alert is-success">
                        Your account has been upgraded successfully!
                    </div>
                }

                {this.props.msg === 'Already upgrade' &&
                    <div className="alert is-error">
                        Already upgraded to enterprise plan!
                    </div>
                }
            </div>
        );
    }
}