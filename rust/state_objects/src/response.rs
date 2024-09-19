//! Provides the state_objects response type;

use interchain_core::error::ErrorMessage;
use interchain_schema::value::ResponseValue;

/// Response is the response type for state_objects methods.
// TODO: constrain R
pub type Response<'a, R, E: ResponseValue = ErrorMessage> = Result<R, <<E as ResponseValue>::MaybeBorrowed<'a> as ToOwned>::Owned>;
